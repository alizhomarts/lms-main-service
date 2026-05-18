-- +goose Up

INSERT INTO courses (id, name, description, created_at, updated_at)
VALUES
    (1, 'Golang Developer', 'A practical backend development course focused on Go, PostgreSQL, REST API design, Docker, testing, and production-ready service architecture.', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Python Developer', 'A hands-on course for learning Python programming, web development fundamentals, working with data, APIs, and backend application structure.', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
    ON CONFLICT (id) DO NOTHING;

INSERT INTO chapters (id, name, description, "order", course_id, created_at, updated_at)
VALUES
    (1, 'Control Structures', 'This chapter explains how to control program flow in Go using conditional statements, loops, switch statements, break, continue, and defer.', 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Working with Functions', 'This chapter introduces functions in Go, including parameters, return values, multiple return values, named returns, and error handling patterns.', 2, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 'Python Basics', 'This chapter introduces Python syntax, variables, data types, conditions, loops, and basic input/output operations.', 1, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
    ON CONFLICT (id) DO NOTHING;

INSERT INTO lessons (id, name, description, content, "order", chapter_id, created_at, updated_at)
VALUES
    (
        1,
        'If-else Statement in Golang',
        'This lesson explains conditional branching in Go using if, else if, and else statements.',
        'In Go, the if statement is used to execute a block of code only when a specific condition is true. Unlike some other programming languages, Go does not require parentheses around the condition. However, curly braces are required. You can also use else if and else blocks to handle multiple execution paths. Go also supports short variable declarations inside if statements, which is commonly used when checking values returned from functions.',
        1,
        1,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        2,
        'Switch Statement in Golang',
        'This lesson explains how to use switch statements for cleaner conditional logic.',
        'The switch statement in Go provides a cleaner way to write multiple conditional branches. A switch can compare values, evaluate expressions, or be used without an expression as an alternative to long if-else chains. In Go, cases do not fall through by default, which makes switch statements safer and easier to understand.',
        2,
        1,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        3,
        'Functions in Go',
        'This lesson explains how to define and use functions in Go.',
        'Functions are one of the core building blocks of Go programs. A function can accept parameters, return one or more values, and help organize code into reusable blocks. Go functions often return an error as the last return value, which makes error handling explicit and easy to follow.',
        1,
        2,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        4,
        'Variables in Python',
        'This lesson explains how variables work in Python.',
        'In Python, variables are used to store values such as numbers, text, lists, dictionaries, and other objects. Python is dynamically typed, which means you do not need to declare the variable type explicitly. A variable is created when a value is assigned to it.',
        1,
        3,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    )
    ON CONFLICT (id) DO NOTHING;

SELECT setval('courses_id_seq', (SELECT MAX(id) FROM courses));
SELECT setval('chapters_id_seq', (SELECT MAX(id) FROM chapters));
SELECT setval('lessons_id_seq', (SELECT MAX(id) FROM lessons));

-- +goose Down

DELETE FROM lessons WHERE id IN (1, 2, 3, 4);
DELETE FROM chapters WHERE id IN (1, 2, 3);
DELETE FROM courses WHERE id IN (1, 2);