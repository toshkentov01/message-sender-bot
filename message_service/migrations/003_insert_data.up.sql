INSERT INTO messages(
    id,
    content,
    priority
) VALUES (uuid_generate_v4(), 'first_message', 'high'), 
    (uuid_generate_v4(), 'second_message', 'low'),
    (uuid_generate_v4(), 'third_message', 'medium'),
    (uuid_generate_v4(), 'last_message', 'low');
