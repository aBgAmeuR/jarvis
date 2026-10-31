import spec from "@jarvis/openapi/openapi.yaml?raw";

import "@scalar/api-reference/style.css";
import { createApiReference } from "@scalar/api-reference";

import "./style.css";

createApiReference("#app", {
  content: spec,
  expandAllParameters: true,
  hideTestRequestButton: true,
});
