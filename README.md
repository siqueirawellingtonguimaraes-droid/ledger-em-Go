# 💰 GoLedger

> Implementação de um pequeno Ledger financeiro em Go, explorando conceitos utilizados em arquiteturas financeiras reais.

---

# 🧠 Sobre o Projeto

GoLedger é um projeto desenvolvido com foco em estudar como sistemas financeiros lidam com consistência, rastreabilidade e integridade de dados.

Ao invés de seguir o modelo tradicional de CRUD, o projeto foi construído explorando conceitos encontrados em sistemas financeiros reais, onde o saldo não é armazenado diretamente, mas calculado a partir do histórico de transações.

Mesmo sendo um projeto relativamente pequeno, ele foi pensado para explorar problemas comuns em domínios financeiros:

- precisão monetária
- consistência de dados
- modelagem de domínio
- imutabilidade
- rastreabilidade de operações
- separação de responsabilidades

O principal objetivo foi aprofundar conhecimentos em engenharia de software aplicada a sistemas críticos, utilizando Go como linguagem principal.

---

# 🎯 Objetivos Técnicos

O projeto foi desenvolvido para explorar:

- modelagem de transações financeiras
- cálculo de saldo baseado em eventos
- arquitetura orientada a domínio
- separação entre domínio e infraestrutura
- consistência de dados financeiros
- redução de acoplamento
- organização de regras de negócio

---

# 🏗️ Conceitos de Engenharia Aplicados

## 💵 Representação Monetária com Inteiros

Valores monetários são representados utilizando inteiros ao invés de floating point.

```txt
1000 = R$10,00
```

Essa abordagem evita problemas de precisão e arredondamento.

---

## 📚 Saldo Baseado em Histórico de Transações

O saldo não é persistido diretamente no banco.

Ele é calculado dinamicamente a partir do histórico de transações.

```txt
+1000
-250
+500
------
1250
```

Esse modelo melhora:

- auditabilidade
- rastreabilidade
- integridade histórica
- previsibilidade de comportamento

---

## 🧩 Organização Inspirada em DDD

A estrutura da aplicação foi organizada utilizando conceitos inspirados em Domain-Driven Design, e busca manter:

- regras de negócio isoladas
- baixo acoplamento
- alta coesão
- facilidade de manutenção e evolução

---

## 🔌 Injeção de Dependências

As dependências são injetadas explicitamente, reduzindo acoplamento entre camadas e melhorando testabilidade da aplicação.

---

## 🧱 Value Objects

O projeto utiliza Value Objects para evitar Primitive Obsession.

Exemplos:

- Money
- Currency (BRL, USD)
- Direction (credit, debit)

Isso torna o domínio mais expressivo, seguro e consistente.

---

# ⚙️ Conceitos Financeiros Explorados

## 💳 Transações como Fonte da Verdade

Toda movimentação financeira é representada como uma transação. O estado do sistema é derivado das transações registradas.

---

## ⚖️ Consistência Acima de Simplicidade

Durante o desenvolvimento, ficou evidente como sistemas financeiros exigem muito mais preocupação com:

- consistência
- previsibilidade
- rastreabilidade
- integridade de dados

do que aplicações CRUD tradicionais.

---

# 🚀 Melhorias Futuras

Mesmo sendo um projeto de estudo, diversos cenários reais surgiram durante o desenvolvimento.

## Planejado

- auditoria de transações
- idempotência
- controle de concorrência
- proteção contra alteração de registros
- triggers para segurança financeira
- processamento orientado a eventos
- reconciliação financeira
- suporte a double-entry bookkeeping

---

# 🧠 O Que Este Projeto Demonstra

Este projeto foi desenvolvido para demonstrar conhecimentos em:

- engenharia de software aplicada a domínios críticos
- arquitetura backend
- modelagem financeira
- Domain-Driven Design
- organização de código
- desacoplamento entre camadas
- consistência de dados
- design orientado a domínio

---

# 📦 Stack

- Go (Golang)
- SQLite (Em um sistema real, nunca se deve usar SQLite para uma aplicação desse porte)
- Domain-Driven Design (DDD)
- Repository Pattern
- Dependency Injection
- Value Objects

---

# 📌 Observação

Este projeto não possui objetivo de se tornar um sistema bancário real.

O foco principal é estudo arquitetural e aprofundamento em conceitos utilizados em sistemas financeiros modernos.

---

# 👨‍💻 Considerações Finais

GoLedger foi um projeto extremamente importante para entender como aplicações financeiras vão muito além de operações básicas de CRUD.

Durante o desenvolvimento, foi possível explorar desafios reais relacionados a:

- precisão monetária
- consistência de dados
- modelagem de domínio
- rastreabilidade de operações
- integridade financeira

Projetos como esse ajudam a desenvolver uma visão muito mais profunda sobre engenharia de software, principalmente quando começamos a trabalhar com domínios onde confiabilidade e consistência são requisitos fundamentais.
