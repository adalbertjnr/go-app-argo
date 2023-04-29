CREATE TABLE produtos (
    id serial primary key,
    nome varchar,
    descricao varchar,
    preco decimal,
    quantidade integer
);

INSERT INTO produtos (nome, descricao, preco, quantidade) values('Bike', 'Red', 1500, 5);

SELECT * FROM produtos;