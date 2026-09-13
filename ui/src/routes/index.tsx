import { createFileRoute } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";

const Route = createFileRoute("/")({
  component: LandingPage,
});

function LandingPage() {
  return (
    <section className="mx-auto flex min-h-screen max-w-4xl items-center px-6 py-16">
      <Card className="w-full">
        <CardHeader>
          <p className="text-sm font-medium text-muted-foreground">Go + React Template</p>
          <CardTitle className="text-4xl">Build your next service.</CardTitle>
          <CardDescription className="text-base">
            A production-ready starting point with Go, Connect RPC, React, and TanStack Query.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <Button>Get started</Button>
        </CardContent>
      </Card>
    </section>
  );
}

export { Route };
