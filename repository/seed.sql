/* Seed a test user (johndoe@example.com/secret) */
INSERT INTO users("email", "password", "first_name", "last_name") VALUES
('johndoe@example.com', '$2y$10$i1jALNbm199VCCNpA0hNSe/3e34WOPdJy1yk7X.GrNAaJqGaUFgMK', 'John', 'Doe');

/* Seed some organizations */
INSERT INTO organizations("name", "email", "phone", "address", "city", "region", "country", "postal_code") VALUES
('Example Ltd.', 'contact@example.com', '555-8451-555', 'Somewhere 17', 'Rainbow', 'Over', 'US', '555555');

/* Seed some contacts */
INSERT INTO contacts("first_name", "last_name", "email", "phone", "address", "city", "region", "country", "postal_code") VALUES
('Johanna', 'Doe', 'johannadoe@example.com', '555-7421-555', 'Somewhere 11', 'Rainbow', 'Over', 'US', '555555');