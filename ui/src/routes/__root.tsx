import { Outlet, createRootRoute } from "@tanstack/react-router";

const Route = createRootRoute({
  component: RootLayout,
});

function RootLayout() {
  return (
    <main className="min-h-screen bg-background text-foreground">
      <Outlet />
    </main>
  );
}

export { Route };
