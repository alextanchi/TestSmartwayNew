CREATE TABLE employees
(
    id              uuid    NOT NULL,
    name            varchar NOT NULL,
    surname         varchar NOT NULL,
    phone           varchar NOT NULL,
    company_id      integer NOT NULL,
    passport_type   varchar NOT NULL,
    passport_number varchar NOT NULL,
    department_id   uuid    NOT NULL,
    CONSTRAINT employees_pkey PRIMARY KEY (id) /* устанавливаем ограничение на столбец id,
                                                   это первичный ключ который будет уникальным*/
    /*прописать форейн кей*/

);

CREATE TABLE department
(
    id          uuid    NOT NULL,
    name        varchar NOT NULL,
    phone       varchar NOT NULL,

    CONSTRAINT department_pkey PRIMARY KEY (id)

);











