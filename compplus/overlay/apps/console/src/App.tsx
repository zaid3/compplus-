import { ConfirmDialog, Toasts } from "@probo/ui";
import { RouterProvider } from "react-router";

import { ThemeManager } from "./components/ThemeManager";
import { router } from "./routes";

export function App() {
  return (
    <>
      <ThemeManager />
      <RouterProvider router={router} />
      <Toasts />
      <ConfirmDialog />
    </>
  );
}
