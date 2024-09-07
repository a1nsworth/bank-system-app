# Bank System App

## Overview

Bank System App is a Go application that manages bank-related data, including banks, bank offices, ATMs, employees, users, payment accounts, and credit accounts. The project utilizes GORM for ORM functionality and SQLite for data storage.

## Features

- Manage banks, bank offices, ATMs, employees, and users.
- Create, read, update, and delete (CRUD) operations for all entities.
- Custom error handling and logging.
- Integration with GORM for ORM functionality.
- Writing tests for CRUD operations.

## Project Structure

- `internal`: Contains internal packages including repositories, services, and models.
- `models`: Defines the data models used by the application.
- `repositories`: Implements CRUD operations for different models.
- `services`: Contains business logic and service layers.
- `tests`: Contains unit and integration tests for the application.

## Setup

### Prerequisites

- Go 1.18 or later
- SQLite (if not using an alternative database)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/a1nsworth/bank-system-app.git
   cd bank-system-app
   
2. Running tests:
   ```bash
   go test ./tests 

# Bank System Data Model

This document describes the data model for a banking system. It provides an overview of the entities and their relationships within the system.

## Entities

### Bank

- **Name**: The name of the bank.
- **Rating**: Rating of the bank (0-100).
- **TotalSum**: Total amount of money held by the bank.
- **InterestRate**: Interest rate offered by the bank (0-20%).

**Relationships:**
- **BankOffices**: One-to-many relationship with BankOffice.
- **BankAtms**: One-to-many relationship with BankAtm.
- **Users**: Many-to-many relationship with User.
- **Employees**: One-to-many relationship with Employee.

### BankOffice

- **Address**: The address of the bank office.
- **Status**: Status of the bank office (e.g., active, able to place ATMs, credit available).
- **Rental**: Rental cost of the bank office.

**Relationships:**
- **Bank**: Many-to-one relationship with Bank.
- **BankAtms**: One-to-many relationship with BankAtm.
- **Employees**: One-to-many relationship with Employee.

### BankAtm

- **Name**: The name of the ATM.
- **Status**: Status of the ATM (e.g., active, has money, working to dispense money).
- **Amortization**: Amortization value of the ATM.

**Relationships:**
- **Bank**: Many-to-one relationship with Bank.
- **BankOffice**: Many-to-one relationship with BankOffice.

### Employee

- **Position**: Position of the employee.
- **Status**: Status of the employee (e.g., remote, can give loans).
- **Salary**: Salary of the employee.

**Relationships:**
- **Bank**: Many-to-one relationship with Bank.
- **BankOffice**: Many-to-one relationship with BankOffice.

### User

- **PlaceOfWork**: The place of work for the user.
- **MonthlyIncome**: Monthly income of the user.
- **BankCreditScore**: Credit score of the user with the bank.

**Relationships:**
- **BanksUsed**: Many-to-many relationship with Bank.
- **CreditAccounts**: One-to-many relationship with CreditAccount.
- **PaymentAccounts**: One-to-many relationship with PaymentAccount.

### PaymentAccount

- **Balance**: Balance of the payment account.

**Relationships:**
- **User**: Many-to-one relationship with User.
- **Bank**: Many-to-one relationship with Bank.

### CreditAccount

- **LoanStartDate**: Start date of the loan.
- **LoanEndDate**: End date of the loan.
- **LoanDurationMonths**: Duration of the loan in months.
- **LoanAmount**: Amount of the loan.
- **MonthlyPayment**: Monthly payment for the loan.
- **InterestRate**: Interest rate for the loan.

**Relationships:**
- **User**: Many-to-one relationship with User.
- **Bank**: Many-to-one relationship with Bank.
- **Employee**: Many-to-one relationship with Employee.
- **PaymentAccount**: Many-to-one relationship with PaymentAccount.

## Status Types

### OfficeStatus

- **OfficeActive**: Indicates if the office is active.
- **OfficeAbleToPlaceAtm**: Indicates if ATMs can be placed at the office.
- **OfficeCreditAvailable**: Indicates if credit is available at the office.

### AtmStatus

- **AtmActive**: Indicates if the ATM is active.
- **AtmHaveMoney**: Indicates if the ATM has money.
- **AtmWorkToDispenseMoney**: Indicates if the ATM is working to dispense money.
- **AtmAbleWithdraw**: Indicates if the ATM is able to withdraw money.

### EmployeeStatus

- **EmployeeIsRemote**: Indicates if the employee is remote.
- **EmployeeCanGiveLoans**: Indicates if the employee can give loans.

## Relationships Overview

- **One-to-Many**: A single entity is related to multiple instances of another entity.
  - Example: A Bank has many BankOffices.
  
- **Many-to-One**: Multiple instances of an entity are related to a single instance of another entity.
  - Example: Many BankAtms are associated with one Bank.

- **Many-to-Many**: Multiple instances of one entity are related to multiple instances of another entity.
  - Example: Users and Banks have a many-to-many relationship.


