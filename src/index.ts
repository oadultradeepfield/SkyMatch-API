import {Container, getContainer} from "@cloudflare/containers";

export class SkyMatchContainer extends Container<Env> {
    defaultPort = 8080;
    sleepAfter = "2m";
}

export default {
    async fetch(request: Request, env: Env): Promise<Response> {
        const container = getContainer(env.SKY_MATCH_CONTAINER);
        return await container.fetch(request);
    },
};
