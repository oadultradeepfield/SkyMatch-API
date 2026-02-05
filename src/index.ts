import {Container, getContainer} from "@cloudflare/containers";
import {env} from "cloudflare:workers";

export class SkyMatchContainer extends Container<Env> {
    defaultPort = 8080;
    sleepAfter = "2m";
    envVars = {
        NOVA_API_KEY: env.NOVA_API_KEY,
        CF_ACCOUNT_ID: env.CLOUDFLARE_ACCOUNT_ID,
        CF_KV_NAMESPACE_ID: env.CLOUDFLARE_KV_NAMESPACE_ID,
        CF_API_TOKEN: env.CLOUDFLARE_API_TOKEN,
    };
}

export default {
    async fetch(request: Request, env: Env): Promise<Response> {
        const container = getContainer(env.SKY_MATCH_CONTAINER);
        return await container.fetch(request);
    },
};
