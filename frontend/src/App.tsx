import { BookManage, BookManageJson, BookState } from "./domain/book.ts";
import { use, useActionState } from "react";
import { addBookAction, searchBookAction } from "./actions/book.ts";

async function fetchManageBook() {
  await new Promise((resolve) => setTimeout(resolve, 1000));
  const response = await fetch("http://localhost:8080/books");
  const data = (await response.json()) as BookManageJson[];
  return data.map((book) => new BookManage(book.id, book.name, book.status));
}

const fetchManageBookPromise = fetchManageBook();

function App() {
  const initialBooks = use(fetchManageBookPromise);
  const [bookState, updateBookState, isPending] = useActionState(
    async (prevState: BookState, formData: FormData) => {
      const formType = formData.get("formType") as string;

      const actionMap = {
        add: () => addBookAction(prevState, formData),
        search: () => searchBookAction(prevState, formData),
      } as const;

      if (formType !== "add" && formType !== "search") {
        throw new Error(`Invalid form type: ${formType}`);
      }

      const action = actionMap[formType];

      return await action();
    },
    {
      books: initialBooks,
    },
  );

  return (
    <>
      <div>
        <form action={updateBookState}>
          <input type="hidden" name="formType" value="add" />
          <input type="text" name="bookName" placeholder="書籍名" />
          <button type="submit" disabled={isPending}>
            追加
          </button>
        </form>
        <form action={updateBookState}>
          <input type="hidden" name="formType" value="search" />
          <input type="text" name="keyword" placeholder="書籍名で検索" />
          <button type="submit" disabled={isPending}>
            検索
          </button>
        </form>
        <div>
          <ul>
            {bookState.books.map((book: BookManage) => {
              return <li key={book.id}>{book.name}</li>;
            })}
          </ul>
        </div>
      </div>
    </>
  );
}

export default App;
