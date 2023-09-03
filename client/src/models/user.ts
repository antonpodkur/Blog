export interface User {
  id: string | null,
  name: string,
  email: string,
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

export interface UserSignIn {
  email: string,
  password: string
}

export interface UserResponse {
  id: string | null,
  name: string,
  email: string,
  role: string,
  created_at: Date,
  updated_at: Date
}

export const mapUserResponseToUser = (response: UserResponse): User => {
  const user: User = {
    ...response,
    createdAt: response.created_at,
    updatedAt: response.updated_at
  }

  return user
}
