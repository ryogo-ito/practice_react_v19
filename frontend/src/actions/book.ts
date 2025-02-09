import { BookManage, BookManageJson, BookState } from "../domain/book.ts";

export const addBookAction = async (
  prevState: BookState,
  formData: FormData,
): Promise<BookState> => {
  const name = formData.get("bookName") as string;

  if (!name) {
    throw new Error("Book name is required");
  }

  const response = await fetch("http://localhost:8080/books", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ name }),
  });

  if (!response.ok) {
    throw new Error("Failed to add book");
  }

  const newBook = await response.json();

  return {
    ...prevState,
    books: [...prevState.books, newBook],
  };
};

export const searchBookAction = async (
  prevState: BookState,
  formData: FormData,
): Promise<BookState> => {
  const keyword = formData.get("keyword") as string;

  if (!keyword) {
    throw new Error("keyword not found");
  }

  const response = await fetch(
    `http://localhost:8080/books?keyword=${keyword}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    },
  );

  if (!response.ok) {
    throw new Error("Failed to search book");
  }

  const books = (await response.json()) as BookManageJson[];

  return {
    ...prevState,
    books: books.map(
      ({ id, name, status }) => new BookManage(id, name, status),
    ),
  };
};

export const updateBookStatusAction = async (
  prevState: BookState,
  formData: FormData,
): Promise<BookState> => {
  const status = formData.get("status");
  const id = formData.get("id");

  if (!status || !id) {
    throw new Error("Status not found");
  }

  const response = await fetch(`http://localhost:8080/books/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ status }),
  });

  if (!response.ok) {
    throw new Error("Failed to update book status");
  }

  const updatedBook = (await response.json()) as BookManageJson;

  const updatedBooks = prevState.books.map((book) =>
    book.id === updatedBook.id
      ? new BookManage(updatedBook.id, updatedBook.name, updatedBook.status)
      : book,
  );

  return {
    ...prevState,
    books: updatedBooks,
  };
};
