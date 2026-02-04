import {Container, getContainer} from "@cloudflare/containers";
import {env} from "cloudflare:workers";

export class SkyMatchContainer extends Container<Env> {
    defaultPort = 8080;
    sleepAfter = "2m";
    envVars = {
        NOVA_API_KEY: env.NOVA_API_KEY,
    };
}

export default {
    async fetch(request: Request, env: Env): Promise<Response> {
        const container = getContainer(env.SKY_MATCH_CONTAINER);
        return await container.fetch(request);
    },
};
