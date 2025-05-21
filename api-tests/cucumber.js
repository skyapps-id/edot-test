module.exports = {
  default: {
    require: ['features/step_definitions/*.js'],
    format: ['progress', 'html:reports/cucumber-report.html'],
    publishQuiet: true,
  },
};
