// js/app.js

const API_URL = "http://localhost:8080";

// ==============================
// HELPERS
// ==============================

function formatCurrency(value) {
    return new Intl.NumberFormat("pt-BR", {
        style: "currency",
        currency: "BRL"
    }).format(value / 100);
}

function showToast(message, type = "success") {
    const toast = document.createElement("div");

    toast.className = `toast ${type}`;
    toast.innerText = message;

    document.body.appendChild(toast);

    setTimeout(() => {
        toast.classList.add("show");
    }, 100);

    setTimeout(() => {
        toast.classList.remove("show");

        setTimeout(() => {
            toast.remove();
        }, 300);
    }, 3000);
}

async function request(url, options = {}) {
    try {
        const response = await fetch(url, {
            headers: {
                "Content-Type": "application/json"
            },
            ...options
        });

        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || "Erro inesperado");
        }

        if (response.status === 204) {
            return null;
        }

        return await response.json();
    } catch (err) {
        console.error(err);
        showToast(err.message, "error");
        throw err;
    }
}

// ==============================
// CREATE ACCOUNT
// ==============================

async function createAccount(event) {
    event.preventDefault();

    const nameInput = document.querySelector("#account-name");
    const currencyInput = document.querySelector("#account-currency");

    const payload = {
        name: nameInput.value,
        currency: currencyInput.value || "BRL"
    };

    await request(`${API_URL}/accounts/`, {
        method: "POST",
        body: JSON.stringify(payload)
    });

    showToast("Conta criada com sucesso");

    nameInput.value = "";
}

// ==============================
// GET ACCOUNT
// ==============================

async function getAccountById() {
    const accountId = document.querySelector("#search-account-id").value;

    if (!accountId) {
        showToast("Informe um ID", "error");
        return;
    }

    const response = await request(
        `${API_URL}/accounts/${accountId}`
    );

    renderAccount(response);
}

function renderAccount(account) {
    const container = document.querySelector("#account-result");

    container.innerHTML = `
        <div class="account-card">
            <div>
                <span class="account-label">Conta</span>
                <h3>${account.name}</h3>
            </div>

            <div>
                <span class="account-label">Moeda</span>
                <h4>${account.currency}</h4>
            </div>

            <div>
                <span class="account-label">Saldo</span>
                <h2>${formatCurrency(account.balance)}</h2>
            </div>
        </div>
    `;
}

// ==============================
// TRANSACTION ENTRIES
// ==============================

function createReceiverField() {
    const container = document.querySelector("#receivers-container");

    const div = document.createElement("div");

    div.className = "receiver-item";

    div.innerHTML = `
        <input 
            type="text"
            class="receiver-account"
            placeholder="Conta destino"
        >

        <input
            type="number"
            class="receiver-amount"
            placeholder="Valor"
        >

        <button class="remove-button">
            Remover
        </button>
    `;

    div.querySelector(".remove-button")
        .addEventListener("click", () => {
            div.remove();
            validateTransactionBalance();
        });

    div.querySelector(".receiver-amount")
        .addEventListener("input", validateTransactionBalance);

    container.appendChild(div);
}

// ==============================
// BALANCE VALIDATION
// ==============================

function validateTransactionBalance() {
    const totalInput =
        document.querySelector("#transaction-total");

    const receiverAmounts =
        document.querySelectorAll(".receiver-amount");

    const total = Number(totalInput.value || 0);

    let distributed = 0;

    receiverAmounts.forEach(input => {
        distributed += Number(input.value || 0);
    });

    const diff = total - distributed;

    const balanceLabel =
        document.querySelector("#balance-validation");

    if (diff === 0) {
        balanceLabel.innerHTML = `
            Balanceado ✓
        `;

        balanceLabel.className = "success-text";
    } else {
        balanceLabel.innerHTML = `
            Diferença: ${diff}
        `;

        balanceLabel.className = "error-text";
    }

    return diff === 0;
}

// ==============================
// CREATE TRANSACTION
// ==============================

async function createTransaction(event) {
    event.preventDefault();

    const isValid = validateTransactionBalance();

    if (!isValid) {
        showToast(
            "Débito e crédito devem ser iguais",
            "error"
        );

        return;
    }

    const description =
        document.querySelector("#transaction-description").value;

    const originAccount =
        document.querySelector("#origin-account").value;

    const total =
        Number(document.querySelector("#transaction-total").value);

    const receiverAccounts =
        document.querySelectorAll(".receiver-account");

    const receiverAmounts =
        document.querySelectorAll(".receiver-amount");

    const entries = [];

    // DEBIT ENTRY
    entries.push({
        account_id: originAccount,
        amount: total,
        currency: "BRL",
        direction: "debit"
    });

    // CREDIT ENTRIES
    receiverAccounts.forEach((accountInput, index) => {

        const amount =
            Number(receiverAmounts[index].value);

        entries.push({
            account_id: accountInput.value,
            amount,
            currency: "BRL",
            direction: "credit"
        });
    });

    const payload = {
        description,
        entries
    };

    await request(`${API_URL}/transactions/`, {
        method: "POST",
        body: JSON.stringify(payload)
    });

    showToast("Transação criada");

    loadTransactions();

    document
        .querySelector("#transaction-form")
        .reset();

    document.querySelector(
        "#receivers-container"
    ).innerHTML = "";

    createReceiverField();
}

// ==============================
// LOAD TRANSACTIONS
// ==============================

async function loadTransactions() {
    const tbody =
        document.querySelector("#transactions-table");

    if (!tbody) return;

    const transactions = await request(
        `${API_URL}/transactions/`
    );

    tbody.innerHTML = "";

    transactions.forEach(transaction => {

        const total = transaction.entries
            .filter(entry => entry.direction === "DEBIT")
            .reduce((acc, item) => acc + item.amount, 0);

        const tr = document.createElement("tr");

        tr.innerHTML = `
            <td>${transaction.id}</td>
            <td>${transaction.description}</td>
            <td>
                <span class="badge success">
                    Concluído
                </span>
            </td>
            <td>${formatCurrency(total)}</td>
        `;

        tbody.appendChild(tr);
    });
}

// ==============================
// GET ENTRIES BY ACCOUNT
// ==============================

async function getEntriesByAccount() {
    const accountId =
        document.querySelector("#entries-account-id").value;

    const tbody =
        document.querySelector("#entries-table");

    const entries = await request(
        `${API_URL}/transactions/entries/${accountId}`
    );

    tbody.innerHTML = "";

    entries.forEach(entry => {

        const tr = document.createElement("tr");

        tr.innerHTML = `
            <td>${entry.id}</td>
            <td>${entry.direction}</td>
            <td>${formatCurrency(entry.amount)}</td>
            <td>${entry.currency}</td>
        `;

        tbody.appendChild(tr);
    });
}

// ==============================
// DASHBOARD METRICS
// ==============================

async function loadDashboardMetrics() {
    const transactions =
        await request(`${API_URL}/transactions/`);

    const totalTransactions =
        transactions.length;

    let totalMoved = 0;

    transactions.forEach(transaction => {

        const debit = transaction.entries
            .find(entry => entry.direction === "DEBIT");

        if (debit) {
            totalMoved += debit.amount;
        }
    });

    const transactionsLabel =
        document.querySelector("#dashboard-transactions");

    const movedLabel =
        document.querySelector("#dashboard-total-moved");

    if (transactionsLabel) {
        transactionsLabel.innerText =
            totalTransactions;
    }

    if (movedLabel) {
        movedLabel.innerText =
            formatCurrency(totalMoved);
    }
}

// ==============================
// GET ALL ACCOUNTS
// ==============================
async function getAllAccounts() {
    return await request(`${API_URL}/accounts/all`);
}

function renderAccounts(accounts) {
    const container = document.querySelector("#account-result");

    if (!container) return;

    container.innerHTML = "";

    let totalBalance = 0;

    accounts.forEach(account => {
        totalBalance += account.balance;

        const card = document.createElement("div");

        card.className = "account-card";

        card.innerHTML = `
    <div>
        <span class="account-label">Conta</span>
        <h3>${account.name}</h3>
        <small>${account.id}</small>
    </div>

    <div style="text-align:right">
        <span class="account-label">Saldo</span>
        <h2>${formatCurrency(account.balance)}</h2>
    </div>
`;

        container.appendChild(card);
    });

    const accountsCounter = document.querySelector("#dashboard-total-accounts");
    const totalBalanceLabel = document.querySelector("#dashboard-total-moved");

    if (accountsCounter) {
        accountsCounter.innerText = accounts.length;
    }

    if (totalBalanceLabel) {
        totalBalanceLabel.innerText = formatCurrency(totalBalance);
    }
}

async function loadAccountsDashboard() {
    try {
        const accounts = await getAllAccounts();
        renderAccounts(accounts);
    } catch (err) {
        console.error("Erro ao carregar contas", err);
    }
}
// ==============================
// INITIALIZE
// ==============================

document.addEventListener("DOMContentLoaded", () => {

    // CREATE ACCOUNT
    const accountForm =
        document.querySelector("#create-account-form");

    if (accountForm) {
        accountForm.addEventListener(
            "submit",
            createAccount
        );
    }

    // SEARCH ACCOUNT
    const searchButton =
        document.querySelector("#search-account-button");

    if (searchButton) {
        searchButton.addEventListener(
            "click",
            getAccountById
        );
    }

    // TRANSACTION FORM
    const transactionForm =
        document.querySelector("#transaction-form");

    if (transactionForm) {
        transactionForm.addEventListener(
            "submit",
            createTransaction
        );
    }

    // ADD RECEIVER
    const addReceiverButton =
        document.querySelector("#add-receiver-button");

    if (addReceiverButton) {
        addReceiverButton.addEventListener(
            "click",
            createReceiverField
        );
    }

    // VALIDATE TOTAL
    const totalInput =
        document.querySelector("#transaction-total");

    if (totalInput) {
        totalInput.addEventListener(
            "input",
            validateTransactionBalance
        );
    }

    // LOAD TRANSACTIONS
    loadTransactions();

    // LOAD DASHBOARD
    loadDashboardMetrics();
    loadAccountsDashboard();

    // INITIAL RECEIVER
    if (document.querySelector("#receivers-container")) {
        createReceiverField();
    }
});