const API_URL = "http://localhost:8080";

let currentAccount = null;
let accounts = [];

async function request(url, options = {}) {

    const response = await fetch(url, {
        headers: {
            "Content-Type": "application/json",
        },
        ...options,
    });

    if (!response.ok) {
        throw new Error("Erro na requisição");
    }

    const contentType = response.headers.get("content-type");

    if (contentType && contentType.includes("application/json")) {
        return await response.json();
    }

    return null;
}

function formatCurrency(value) {

    return new Intl.NumberFormat("pt-BR", {
        style: "currency",
        currency: "BRL",
    }).format(value / 100);

}

async function loadAccounts() {

    try {

        accounts = await request(
            `${API_URL}/accounts/all`
        );

        renderAccountsSelector();

        if (accounts.length > 0) {
            selectAccount(accounts[0].id);
        }

    } catch (err) {
        console.error(err);
    }

}

function renderAccountsSelector() {

    const selector = document.querySelector(
        "#account-selector"
    );

    if (!selector) return;

    selector.innerHTML = "";

    accounts.forEach(account => {

        const option = document.createElement(
            "option"
        );

        option.value = account.id;
        option.textContent = account.name;

        selector.appendChild(option);

    });

}

function selectAccount(accountID) {

    currentAccount = accounts.find(
        account => account.id === accountID
    );

    if (!currentAccount) return;

    document.querySelector(
        "#account-name"
    ).innerText = currentAccount.name;

    document.querySelector(
        "#account-id"
    ).innerText = currentAccount.id;

    document.querySelector(
        "#account-balance"
    ).innerText = formatCurrency(
        currentAccount.balance
    );

    loadTransactions(currentAccount.id);

}

async function loadTransactions(accountID) {

    try {

        const entries = await request(
            `${API_URL}/transactions/entries/${accountID}`
        );

        renderTransactions(entries);

    } catch (err) {
        console.error(err);
    }

}

function renderTransactions(entries) {

    const table = document.querySelector(
        "#transactions-table"
    );

    const totalTransfers = document.querySelector(
        "#total-transfers"
    );

    if (!table) return;

    table.innerHTML = "";

    totalTransfers.innerText = entries.length;

    if (entries.length === 0) {

        table.innerHTML = `
            <tr>
                <td colspan="4">
                    Nenhuma transferência encontrada
                </td>
            </tr>
        `;

        return;
    }

    entries.forEach(entry => {

        const row = document.createElement("tr");

        row.innerHTML = `
            <td>
                Transferência
            </td>

            <td>
                ${entry.account_id}
            </td>

            <td>
                ${formatCurrency(entry.amount)}
            </td>

            <td>
                <span class="
    status
    ${entry.direction === 'debit'
                ? 'danger'
                : 'success'}
">
    ${entry.direction}
</span>
            </td>
        `;

        table.appendChild(row);

    });

}

async function createAccount(event) {

    event.preventDefault();

    try {

        const payload = {
            name: document.querySelector(
                "#account-create-name"
            ).value,

            currency: "BRL",
        };

        await request(
            `${API_URL}/accounts/`,
            {
                method: "POST",
                body: JSON.stringify(payload),
            }
        );

        closeModal();

        document.querySelector(
            "#create-account-form"
        ).reset();

        await loadAccounts();

    } catch (err) {
        console.error(err);
    }

}

async function createTransfer(event) {

    event.preventDefault();

    if (!currentAccount) return;

    try {

        const amount = getRawCurrencyValue(
            document.querySelector(
                "#transaction-amount"
            ).value
        );

        const destination = document.querySelector(
            "#destination-account"
        ).value;

        const description = document.querySelector(
            "#transaction-description"
        ).value;

        const payload = {
            description,

            entries: [
                {
                    account_id: currentAccount.id,
                    amount: amount,
                    currency: "BRL",
                    direction: "debit",
                },

                {
                    account_id: destination,
                    amount: amount,
                    currency: "BRL",
                    direction: "credit",
                },
            ],
        };

        await request(
            `${API_URL}/transactions/`,
            {
                method: "POST",
                body: JSON.stringify(payload),
            }
        );

        document.querySelector(
            "#transfer-form"
        ).reset();

        await loadAccounts();

        selectAccount(currentAccount.id);

    } catch (err) {
        console.error(err);
    }

}

function copyAccountID() {

    const accountID = document.querySelector(
        "#account-id"
    ).innerText;

    navigator.clipboard.writeText(accountID);

}

function openModal() {

    document.querySelector(
        "#account-modal"
    ).classList.add("active");

}

function closeModal() {

    document.querySelector(
        "#account-modal"
    ).classList.remove("active");

}

function formatInputCurrency(value) {

    value = value.replace(/\D/g, "");

    value = (Number(value) / 100)
        .toFixed(2)
        .replace(".", ",");

    value = value.replace(
        /\B(?=(\d{3})+(?!\d))/g,
        "."
    );

    return `R$ ${value}`;
}

function getRawCurrencyValue(value) {

    return Number(
        value
            .replace(/\D/g, "")
    );

}

function setupEvents() {

    document
        .querySelector("#account-selector")
        .addEventListener(
            "change",
            (event) => {
                selectAccount(event.target.value);
            }
        );

    document
        .querySelector("#copy-account-id")
        .addEventListener(
            "click",
            copyAccountID
        );

    document
        .querySelector("#create-account-form")
        .addEventListener(
            "submit",
            createAccount
        );

    document
        .querySelector("#transfer-form")
        .addEventListener(
            "submit",
            createTransfer
        );

    document
        .querySelector("#open-account-modal")
        .addEventListener(
            "click",
            openModal
        );

    document
        .querySelector("#close-account-modal")
        .addEventListener(
            "click",
            closeModal
        );

    const amountInput = document.querySelector(
        "#transaction-amount"
    );

    amountInput.addEventListener("input", (event) => {

        event.target.value = formatInputCurrency(
            event.target.value
        );

    });

}

window.addEventListener(
    "DOMContentLoaded",
    async () => {

        setupEvents();

        await loadAccounts();

    }
);