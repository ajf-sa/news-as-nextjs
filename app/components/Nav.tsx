
import Link from 'next/link'


const Nav = ({ tags }) => {

    return (
        <>
            <nav className="w-full py-4 border-t border-b bg-gray-100">
                <div>
                    <div className="w-full container mx-auto flex flex-row-reverse sm:flex-row items-center justify-center text-sm font-bold uppercase mt-0 px-6 py-2">
                        {tags.map(tag => (
                            <Link href={`/c/${tag.slug}`} key={tag.id}>
                                <a className="hover:bg-gray-400 rounded py-1 px-3">{tag.name}</a>
                            </Link>
                        ))}

                    </div>
                </div>
            </nav>
        </>
    )
}

export default Nav




