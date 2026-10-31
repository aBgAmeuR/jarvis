import { Input as BaseInput } from "@base-ui/react/input";
import { cn } from "cn";

import { fieldStyles, type Styled } from "../lib/styles.ts";

export type InputProps = Styled<BaseInput.Props>;

/** A one-line text input. Put it in a `Field` to get a label and an error. */
export function Input({ className, ...props }: InputProps) {
  return <BaseInput className={cn(fieldStyles, className)} {...props} />;
}
