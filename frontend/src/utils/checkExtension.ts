import { allowedExtensions } from "@/constants/allowedExtensions";
import path from "path";


export const isAllowed = (file: string) =>
  allowedExtensions.includes(path.extname(file).toLowerCase())

