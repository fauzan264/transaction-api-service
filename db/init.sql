DO $$
BEGIN
    -- Memeriksa dan membuat database
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'transaction_api') THEN
        PERFORM pg_catalog.create_database('transaction_api');
    END IF;

    -- Memeriksa dan membuat schema users
    IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_namespace WHERE nspname = 'users') THEN
        EXECUTE 'CREATE SCHEMA users';
    END IF;

    -- Memeriksa dan membuat schema transactions
    IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_namespace WHERE nspname = 'transactions') THEN
        EXECUTE 'CREATE SCHEMA transactions';
    END IF;

    -- Memeriksa dan membuat tabel users.users
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'users' AND table_name = 'users') THEN
        EXECUTE 'CREATE TABLE users.users (
            id CHAR(36) PRIMARY KEY NOT NULL,
            name VARCHAR(100) NOT NULL,
            nik VARCHAR(16) NOT NULL,
            no_hp VARCHAR(13) NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        )';
    END IF;

    -- Memeriksa dan membuat tabel users.balances
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'users' AND table_name = 'balances') THEN
        EXECUTE 'CREATE TABLE users.balances (
            id CHAR(36) PRIMARY KEY NOT NULL,
            user_id CHAR(36) NOT NULL,
            number VARCHAR(20) NOT NULL,
            balance INT DEFAULT 0 CHECK (balance >= 0),
            updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
        )';
    END IF;

    -- Memeriksa dan membuat tabel transactions.transactions
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'transactions' AND table_name = 'transactions') THEN
        EXECUTE 'CREATE TABLE transactions.transactions (
            id CHAR(36) PRIMARY KEY NOT NULL,
            user_balance_id CHAR(36) NOT NULL,
            amount INT CHECK (amount >= 0),
            status VARCHAR(20) NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
        )';
    END IF;

    -- Memeriksa dan menambahkan constraint foreign key untuk transaksi
    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints
                   WHERE constraint_name = 'fk_transactions_user_balances' 
                     AND table_schema = 'transactions' 
                     AND table_name = 'transactions') THEN
        EXECUTE 'ALTER TABLE transactions.transactions
                 ADD CONSTRAINT fk_transactions_user_balances
                 FOREIGN KEY (user_balance_id) REFERENCES users.balances(id)';
    END IF;

    -- Memeriksa dan menambahkan constraint foreign key untuk balance dan users
    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints
                   WHERE constraint_name = 'fk_user_balances_users' 
                     AND table_schema = 'users' 
                     AND table_name = 'balances') THEN
        EXECUTE 'ALTER TABLE users.balances
                 ADD CONSTRAINT fk_user_balances_users
                 FOREIGN KEY (user_id) REFERENCES users.users(id)';
    END IF;

END
$$;
