import { serve } from "@hono/node-server";
import { Hono } from "hono";
import { cors } from "hono/cors";

const app = new Hono();

type BookManger = {
  id: number;
  name: string;
  status: string;
};

const bookManager: BookManger[] = [
  { id: 1, name: "React入門", status: "在庫あり" },
  { id: 2, name: "TypeScript入門", status: "貸出中" },
  { id: 3, name: "Next.js入門", status: "返却済" },
];

app.use(
  "/*",
  cors({
    origin: ["http://localhost:5173"],
    allowMethods: ["GET", "POST", "PUT", "DELETE"],
    allowHeaders: ["Content-Type", "Authorization"],
    exposeHeaders: ["Content-Length"],
    maxAge: 3600,
    credentials: true,
  }),
);

app.put("/books/:id", async (c) => {
  const id = c.req.param("id");
  const body = await c.req.json();
  const status = body.status;

  const book = bookManager.find((book) => book.id === Number(id));

  if (!book) {
    return c.json({ error: "書籍が見つかりません" });
  }

  book.status = status;
  return c.json(book);
});

const PORT = 8080;
console.log(`Server is running on http://localhost:${PORT}`);

serve({
  fetch: app.fetch,
  port: PORT,
});
