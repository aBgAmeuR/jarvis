import { createFileRoute } from "@tanstack/react-router";

import { CertificatesPage } from "@/pages/certificates";

export const Route = createFileRoute("/certificates")({
  component: CertificatesPage,
});
