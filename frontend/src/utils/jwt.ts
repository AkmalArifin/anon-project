export function parseJwt(token: string | null): {
    exp: number,
    iat: number,
    id: number,
    username: string
} {
  if (token === null) {
    return {
      exp: 0,
      iat: 0,
      id: 0,
      username: ""
    }
  } else {
    return JSON.parse(atob(token.split(".")[1]))
  }
}