import { PostType } from "lib/interface"
import Link from 'next/link'
import Image from 'next/image'
import { DateNow } from "./dateNow"
const Post = ({ id, title, description, tags, created_at }: PostType) => {

    return (<>

        <article className="relative flex flex-col text-right shadow my-4  rounded-lg">
            <Link href={`/post/${id}`}>
                <a className="hover:opacity-75 overflow-hidden">
                    <Image
                        className=" py-3 transform hover:scale-110 duration-200"
                        src={`https://source.unsplash.com/collection/1346951/1000x900?sig=` + id}
                        alt=""
                        width={1000}
                        height={900}
                    />
                </a>
            </Link>
            <div className="absolute left-0 z-50 w-full bg-cyan-600">
            <h1 className=" text-white  text-xl  text-center items-center font-bold m-3 pb-1"><span className="text-sm">snap : </span> ajf.sa</h1>
            </div>
            <div className=" absolute bottom-0 z-50 bg-cyan-600">
            <h1 className=" text-white  text-center items-center font-bold m-3 pb-1">
            {tags.map(c => (
                    <Link href={`/c/${c.slug}`} key={c.id?.toString()}>
                        <a className=" text-white text-bold font-bold ">{c.name}</a>
                    </Link>
                ))}
            </h1>
            </div>
            <div className="bg-white flex flex-col justify-start px-4 pb-4">
                
                <Link href={`/post/${id}`}>
                    <a className="text-2xl text-center lg:text-right font-bold hover:text-gray-700 py-4">{title}</a>
                </Link>
                <p className=" text-sm pb-3">
                    نشر في <DateNow date={created_at} />
                </p>
                <Link href={`/post/${id}`}>
                    <a className="pb-10 lg:pb-12 text-lg font-normal text-center lg:text-justify">
                        {description}
                    </a>
                </Link>
                {/* <a  className="uppercase text-gray-800 hover:text-black pb-3">  لكي تعرف اكثر...
                    <span className="font-bold"> <span className="lowercase font-normal"> snap:</span> ajf.sa</span></a> */}
            </div>
        </article>
    </>)
}

export default Post