CREATE TABLE schools (
  id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name       text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE users (
  id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  firebase_uid text NOT NULL UNIQUE,
  email        text NOT NULL,
  display_name text,
  created_at   timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE user_school_roles (
  id        uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id   uuid NOT NULL REFERENCES users(id)   ON DELETE CASCADE,
  school_id uuid NOT NULL REFERENCES schools(id) ON DELETE CASCADE,
  role      text NOT NULL CHECK (role IN ('parent','teacher','admin','driver')),
  UNIQUE (user_id, school_id)          -- one role per user per school
);
CREATE INDEX ON user_school_roles (school_id);

CREATE TABLE students (
  id               uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  school_id        uuid NOT NULL REFERENCES schools(id) ON DELETE CASCADE,
  full_name        text NOT NULL,
  admission_number text,
  created_at       timestamptz NOT NULL DEFAULT now(),
  UNIQUE (id, school_id),                      -- target for the composite FK below
  UNIQUE (school_id, admission_number)         -- NULLs allowed, duplicates are not
);

CREATE TABLE parent_child_links (
  id        uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  school_id uuid NOT NULL,
  parent_id uuid NOT NULL,
  child_id  uuid NOT NULL,
  UNIQUE (parent_id, child_id),
  FOREIGN KEY (child_id, school_id)  REFERENCES students(id, school_id)               ON DELETE CASCADE,
  FOREIGN KEY (parent_id, school_id) REFERENCES user_school_roles(user_id, school_id) ON DELETE CASCADE
);
CREATE INDEX ON parent_child_links (school_id, parent_id);
CREATE INDEX ON parent_child_links (child_id);