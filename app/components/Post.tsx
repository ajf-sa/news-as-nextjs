import { PostType } from "lib/interface"

const Post = ({title,description,category}:PostType)=>{

    return (<>
    
    <article className="flex flex-col text-right shadow my-4 rounded-lg">
              
                <a href="#" className="hover:opacity-75 overflow-hidden ">
                    <img className=" py-3 transform hover:scale-110 duration-200" src="https://source.unsplash.com/collection/1346951/1000x500?sig=1"/>
                </a>
                    <h1 className=" aref text-center font-medium">snap: ajf.sa</h1>
                <div className="bg-white flex flex-col justify-start px-4">
                    <a href="#" className="aref text-cyan-700 text-sm font-bold -pb-4">{category}</a>
                    <a href="#" className="text-2xl text-center lg:text-right font-bold hover:text-gray-700 pb-4">{title}</a>
                    <p  className="aref text-sm pb-3">
                      تاريخ النشر في 
                    </p>
                    <a href="#" className="pb-3 font-bold text-center lg:text-justify">
                       {description}
                         </a>
                    <a href="#" className="uppercase text-gray-800 hover:text-black pb-3">  لكي تعرف اكثر...
                    <span className="font-bold"> <span className="lowercase font-normal"> snap:</span> ajf.sa</span></a>
                </div>
            </article>
    </>)
}

export default Post