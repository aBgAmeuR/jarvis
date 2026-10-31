import { defineConfig } from "orval";

export default defineConfig({
  jarvis: {
    input: "../openapi/openapi.yaml",
    output: {
      target: "./src/gen/jarvis.ts",
      schemas: "./src/gen/model",
      client: "react-query",
      httpClient: "fetch",
      baseUrl: { getBaseUrlFromSpecification: true },
      clean: true,
      override: { fetch: { includeHttpResponseReturnType: false, forceSuccessResponse: true } },
    },
  },
});
