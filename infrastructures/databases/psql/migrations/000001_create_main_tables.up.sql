CREATE TABLE "companies" (
  "id" uuid PRIMARY KEY,
  "name" varchar,
  "created_at" timestamptz,
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "employees" (
  "id" uuid PRIMARY KEY,
  "company_id" uuid,
  "name" varchar,
  "email" varchar,
  "password" varchar,
  "created_at" timestamptz,
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "attendance_logs" (
  "id" integer PRIMARY KEY,
  "employee_id" uuid,
  "note" text,
  "clocked_in_at" timestamptz,
  "clocked_out_at" timestamptz,
  "date" date
);

CREATE INDEX "email_idx" ON "employees" ("email");

ALTER TABLE "employees" ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id");

ALTER TABLE "attendance_logs" ADD FOREIGN KEY ("employee_id") REFERENCES "employees" ("id");
