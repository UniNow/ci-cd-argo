describe('Navigation Test', () => {
  it('should visit home page', () => {
    cy.visit('/')
  })

  it('should visit users page', () => {
    cy.visit('/users')
  })

  it('should visit contact page', () => {
    cy.visit('/contact')
  })
})
