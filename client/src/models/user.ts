export interface User {
  id: string | null,
  email: string,
  name: string,
  role: string,
  createdAt: Date,
  updatedAt: Date
}

export interface UserSignUp {
  name: string,
  email: string,
  password: string,
  passwordConfirm: string
}
