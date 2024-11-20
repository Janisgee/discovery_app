'use client';

import Link from 'next/link';

export default function App() {
  return (
    <div className='p-20 text-center'>
      <div className='text-right'>
        <Button useFor='Sign Up' link='/dashboard/signup' color='btn-grey' />
      </div>
      <div className='my-20 '>
        <h1>Discover Your Side!</h1>
      </div>
      <Button useFor='Login' link='/dashboard/home' color='btn-violet' />
    </div>
  );
}

function Button({ useFor, link, color }) {
  return (
    <div>
      <Link href={link}>
        <button className={color}>{useFor}</button>
      </Link>
    </div>
  );
}

// function Counter() {
//   const [count, setCount] = useState(0);

//   return (
//     <div>
//       <p>You clicked me {count} times</p>
//       <button className='btn-primary' onClick={() => setCount(count + 1)}>
//         Click Me
//       </button>
//     </div>
//   );
// }
