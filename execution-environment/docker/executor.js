try {
    const userCode = process.env.CODE;
    eval(userCode); // Execute the user's code
} catch (err) {
    console.error(`Error: ${err.message}`);
}
