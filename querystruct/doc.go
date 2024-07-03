package querystruct

// Consideration:
// - Run time VS compile time render
// - Composability, reusability, and testability
// - Inversion of control is required for composability
// - Run-time metadata injection
// - Component:
//   - Construct: deal with the SQL template
//   - Composer: compose Construct with QueryModel to render the query
//   - Reference: criteria of reference, it should work by itself

// TODO: Have a robust error handling.
// Error to handle:
//   - Expected QueryModel error should contain which template contain error
//   - If reference is missing it should tell what to do

// TODO: Define unit testing framework. Test should apply to individual Construct
