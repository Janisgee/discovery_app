import Link from 'next/link';

export function Button({ useFor, link, color }) {
  return (
    <div>
      <Link href={link}>
        <button className={color}>{useFor}</button>
      </Link>
    </div>
  );
}
