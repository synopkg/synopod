/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { CommitContext, User, WorkspaceConfig } from "@gitpod/gitpod-protocol";
import { TraceContext } from "@gitpod/gitpod-protocol/lib/util/tracing";
import { inject, injectable } from "inversify";
import { HostContextProvider } from "../auth/host-context-provider";
import { FileProvider } from "../repohost";
import { ConfigInferrer } from "./config-inferrer";

@injectable()
export class ConfigurationService {
    @inject(HostContextProvider) protected readonly hostContextProvider: HostContextProvider;

    // a static cache used to prefetch inferrer related files in parallel in advance
    private requestedPaths = new Set<string>();

    async guessRepositoryConfiguration(
        ctx: TraceContext,
        user: User,
        context: CommitContext,
    ): Promise<string | undefined> {
        const { fileProvider, commitContext } = await this.getRepositoryFileProviderAndCommitContext(
            ctx,
            user,
            context,
        );
        const cache: { [path: string]: Promise<string | undefined> } = {};
        const readFile = async (path: string) => {
            if (path in cache) {
                return await cache[path];
            }
            this.requestedPaths.add(path);
            const content = fileProvider.getFileContent(commitContext, user, path);
            cache[path] = content;
            return await content;
        };
        // eagerly fetch for all files that the inferrer usually asks for.
        this.requestedPaths.forEach((path) => !(path in cache) && readFile(path));
        const configInferrer = new ConfigInferrer();
        const config: WorkspaceConfig = await configInferrer.getConfig({
            // TODO(se) pass down information about currently used IDE. Defaulting to disabling vscode extensions for now, to not bother non VS Code users.
            excludeVsCodeConfig: true,
            config: {},
            read: readFile,
            exists: async (path: string) => !!(await readFile(path)),
        });
        if (!config.tasks) {
            return;
        }
        const configString = configInferrer.toYaml(config);
        return `# This configuration file was automatically generated by Gitpod.
# Please adjust to your needs (see https://www.gitpod.io/docs/introduction/learn-gitpod/gitpod-yaml)
# and commit this file to your remote git repository to share the goodness with others.

# Learn more from ready-to-use templates: https://www.gitpod.io/docs/introduction/getting-started/quickstart

${configString}
`;
    }

    async fetchRepositoryConfiguration(
        ctx: TraceContext,
        user: User,
        context: CommitContext,
    ): Promise<string | undefined> {
        const { fileProvider, commitContext } = await this.getRepositoryFileProviderAndCommitContext(
            ctx,
            user,
            context,
        );
        const configString = await fileProvider.getGitpodFileContent(commitContext, user);
        return configString;
    }

    protected async getRepositoryFileProviderAndCommitContext(
        ctx: TraceContext,
        user: User,
        commitContext: CommitContext,
    ): Promise<{ fileProvider: FileProvider; commitContext: CommitContext }> {
        const { host } = commitContext.repository;
        const hostContext = this.hostContextProvider.get(host);
        if (!hostContext || !hostContext.services) {
            throw new Error(`Cannot fetch repository configuration for host: ${host}`);
        }
        const fileProvider = hostContext.services.fileProvider;
        return { fileProvider, commitContext };
    }
}
