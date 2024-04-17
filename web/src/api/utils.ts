export const baseUrlApi = (url: string) =>
  process.env.NODE_ENV === "development"
    ? `/api/${url}`
    : `https://ams-server.fly.dev/api/${url}`;
export const v1BaseUrlApi = (url: string) =>
  process.env.NODE_ENV === "development"
    ? `/api/v1/${url}`
    : `https://ams-server.fly.dev/api/v1/${url}`;
