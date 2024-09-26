export function parseJwt(token: string): {
    exp: number,
    iat: number,
    id: number,
    username: string
} {
  return JSON.parse(atob(token.split(".")[1]))
}