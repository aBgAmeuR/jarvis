import config from "@jarvis/oxfmt-config";
import { defineConfig } from "oxfmt";

export default defineConfig({ ...config, ignorePatterns: ["src/gen"] });
