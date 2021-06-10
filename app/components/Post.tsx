import { PostType } from 'lib/interface';
import Link from 'next/link';
import Image from 'next/image';
import { DateNow } from './dateNow';

const Post: React.FC<{ post: PostType }> = (props) => {
  const { post } = props;
  return (
    <>
      <article className="relative flex flex-col text-right shadow my-4  rounded-lg">
        <Link href={`/post/${post.slug}`}>
          <a className="hover:opacity-75 overflow-hidden">
            <Image
              className=" py-3 transform hover:scale-110 duration-200"
              src={`${post.feature_image}`}
              alt=""
              width={1000}
              height={1000}
            />
          </a>
        </Link>
        <div className="absolute left-0 z-50 w-full bg-cyan-600">
          <h1 className=" text-white  text-xl  text-center items-center font-bold m-3 pb-1">
            <span className="text-sm">snap : </span> ajf.sa
          </h1>
        </div>
        <div className=" absolute bottom-0 z-50 bg-cyan-600">
          <h1 className=" text-white  text-center items-center font-bold m-3 pb-1"></h1>
        </div>
        <div className="bg-white flex flex-col justify-start px-4 pb-4">
          <Link href={`/post/${post.slug}`}>
            <a className="text-2xl text-center lg:text-right font-bold hover:text-gray-700 py-4">
              {post.title}
            </a>
          </Link>
          <p className=" text-sm pb-3">
            نشر في <DateNow date={post.created_at} />
          </p>

          <p
            className="pb-10 lg:pb-12 text-lg font-normal text-center lg:text-justify"
            dangerouslySetInnerHTML={{ __html: post.html }}
          ></p>

          {/* <a  className="uppercase text-gray-800 hover:text-black pb-3">  لكي تعرف اكثر...
                    <span className="font-bold"> <span className="lowercase font-normal"> snap:</span> ajf.sa</span></a> */}
        </div>
      </article>
    </>
  );
};

export default Post;
