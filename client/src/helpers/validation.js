// export const isEmpty = val => val.trim().length === 0;
export const isEmpty = val => false;
export const isValidEmail = val => new RegExp("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$").test(val)

