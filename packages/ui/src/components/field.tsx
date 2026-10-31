import { Field as BaseField } from "@base-ui/react/field";
import { cn } from "cn";

import type { Styled } from "../lib/styles.ts";

export type FieldProps = Styled<BaseField.Root.Props>;

/** Groups a label, a control, a description and an error, and wires their ids. */
export function Field({ className, ...props }: FieldProps) {
  return <BaseField.Root className={cn("flex min-w-0 flex-col gap-1", className)} {...props} />;
}

export function FieldLabel({ className, ...props }: Styled<BaseField.Label.Props>) {
  return (
    <BaseField.Label
      className={cn("text-xs font-medium text-muted-foreground", className)}
      {...props}
    />
  );
}

export function FieldDescription({ className, ...props }: Styled<BaseField.Description.Props>) {
  return (
    <BaseField.Description className={cn("text-xs text-muted-foreground", className)} {...props} />
  );
}

/** Shown when the control is invalid. */
export function FieldError({ className, ...props }: Styled<BaseField.Error.Props>) {
  return <BaseField.Error className={cn("text-xs text-destructive", className)} {...props} />;
}
