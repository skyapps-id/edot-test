module.exports = {
  default: {
    require: ['features/support/*.js'],
    format: ['progress', 'html:reports/cucumber-report.html'],
    publishQuiet: true,
  },
};
