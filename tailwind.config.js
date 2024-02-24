/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./web-func/**/*.{html,js,templ}"],
    theme: {
        extend: {},
    },
    plugins: [require("daisyui")],
    daisyui: {
        themes: ["light", "dark"],
    },
}