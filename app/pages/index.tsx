import Loyout from 'components/layouts/Layouts'
import { DateNow } from 'components/dateNow'
import Post from 'components/Post'
import { PostType } from 'lib/interface'
const Home = () => {

    const post: PostType = {
        title: "تحكم في التوتر مع Apollo Neuro القابل للارتداء",
        description: "ابتكرت شركة Apollo Neuro واحدًا من أول الأجهزة القابلة للارتداء التي يمكن أن تساعد في تحسين مرونة جسمك في مواجهة",
        category: "تقنية",
    }

    return (
        <>
            <Loyout>
                <section className="w-full flex flex-col items-center px-3">
                    <Post
                        title={post.title}
                        description={post.description}
                        category={post.category}
                    />

                    <div className="flex items-center py-8">
                        <a href="#" className="h-10 w-10 bg-cyan-800 hover:bg-cyan-600 font-semibold text-white text-sm flex items-center justify-center ml-3">1</a>
                        <a href="#" className="h-10 w-10 font-semibold text-gray-800 hover:bg-cyan-600 hover:text-white text-sm flex items-center justify-center ml-3">2</a>
                        <a href="#" className="h-10 w-10 font-semibold text-gray-800 hover:text-gray-900 text-sm flex items-center justify-center ml-3">التالي <i className="fas fa-arrow-right ml-2"></i></a>
                    </div>

                </section>


            </Loyout>
        </>
    )
}

export default Home