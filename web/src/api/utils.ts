export const baseUrlApi = (url: string) =>
  process.env.NODE_ENV === "development"
    ? `/api/${url}`
    : `https://api.ycitoj.top/api/${url}`;
export const v1BaseUrlApi = (url: string) =>
  process.env.NODE_ENV === "development"
    ? `/api/v1/${url}`
    : `https://api.ycitoj.top/api/v1/${url}`;
