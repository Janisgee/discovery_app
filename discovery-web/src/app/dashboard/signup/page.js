import { Button } from '@/app/ui/buttons';

export default async function Signup() {
  return (
    <div className='block-center p-20'>
      <div className='w-full max-w-sm'>
        <div className='mb-10 text-right'>
          <Button useFor='Back' link='/' color='btn-grey' />
        </div>
        <h2 className='mb-10 text-center text-gray-500'>Sign Up</h2>
        <form className='bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4'>
          <div className='mb-4'>
            <label
              className='block text-gray-700 text-sm font-bold mb-2'
              htmlFor='username'
            >
              Username
            </label>
            <input
              className='shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline'
              id='username'
              type='text'
              placeholder='Username'
            />
          </div>
          <div className='mb-4'>
            <label
              className='block text-gray-700 text-sm font-bold mb-2'
              htmlFor='email'
            >
              Email
            </label>
            <input
              className='shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline'
              id='email'
              type='text'
              placeholder='Email'
            />
          </div>
          <div className='mb-6'>
            <label
              className='block text-gray-700 text-sm font-bold mb-2'
              htmlFor='password'
            >
              Password
            </label>
            <input
              className='shadow appearance-none border border-red-500 rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline'
              id='password'
              type='password'
              placeholder='******************'
            />
            <p className='text-red-500 text-xs italic'>
              Please choose a password.
            </p>
          </div>
          <div className='flex items-center justify-between'>
            <button className='btn-violet' type='button'>
              Sign Up
            </button>
            <a
              className='inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800'
              href='#'
            >
              Forgot Password?
            </a>
          </div>
        </form>
        <p className='text-center text-gray-500 text-xs'>
          &copy;2024 Discovery App. All rights reserved.
        </p>
      </div>
    </div>
  );
}

{
  /* // <div>
    //   <h1>Sign up Page</h1>
    //   <form>
    //     <p>Sign Up page</p>

    //     <div>
    //       <label htmlFor='firstName'>First Name</label>
    //       <input
    //         id='firstName'
    //         name='firstName'
    //         type='text'
    //         placeholder='Janis'
    //         autoComplete='firstName'
    //       ></input>
    //     </div>
    //   </form>
    // </div> */
}
