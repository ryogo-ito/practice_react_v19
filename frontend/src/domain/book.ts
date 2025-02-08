export type BookManageJson = {
  id: number;
  name: string;
  status: string;
};

export interface BookState {
  books: BookManage[];
}

export class BookManage {
  constructor(
    public id: number,
    public name: string,
    public status: string,
  ) {}
}
