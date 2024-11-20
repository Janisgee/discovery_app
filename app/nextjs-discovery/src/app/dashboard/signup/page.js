export default async function Signup() {
  return (
    <div>
      <h1>Sign up Page</h1>
      <form>
        <p>Sign Up page</p>

        <div>
          <label htmlFor='firstName'>First Name</label>
          <input
            id='firstName'
            name='firstName'
            type='text'
            placeholder='Janis'
            autoComplete='firstName'
          ></input>
        </div>
      </form>
    </div>
  );
}
