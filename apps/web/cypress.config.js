const { defineConfig } = require("cypress");

module.exports = defineConfig({
  projectId: "7b5hjk",
  e2e: {
    baseUrl: "http://localhost:1323",
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },
});
