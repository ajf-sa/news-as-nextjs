import Loyout from 'components/layouts/Layouts'
import { DateNow } from 'components/dateNow'
export default function Home() {

    

    return (
        <>
            <Loyout>

   
        <section className="w-full flex flex-col items-center px-3">
            <article className="flex flex-col text-right shadow my-4 rounded-lg">
              
                <a href="#" className="hover:opacity-75 overflow-hidden ">
                    <img className=" transform hover:scale-110 duration-200" src="https://source.unsplash.com/collection/1346951/1000x500?sig=1"/>
                </a>
                    <h1 className="text-center font-medium">snap: ajf.sa</h1>
                <div className="bg-white flex flex-col justify-start px-4">
                    <a href="#" className="text-blue-700 text-sm font-bold -pb-4">تقنية</a>
                    <a href="#" className="text-2xl text-center lg:text-right font-bold hover:text-gray-700 pb-4">هاتف ايفون13 قريبا في الاسواق!</a>
                    <p  className="text-sm pb-3">
                      تاريخ النشر في <DateNow   />
                    </p>
                    <a href="#" className="pb-3">
                        تشير التوقعات عن قرب اصدار ايفون 13 في الاسواق مع اقتراب موعد المؤتمر السنوي لشركة ابل...
                    </a>
                    <a href="#" className="uppercase text-gray-800 hover:text-black pb-3">  لكي تعرف اكثر...
                    <span className="font-bold"> <span className="lowercase font-normal"> snap:</span> ajf.sa</span></a>
                </div>
            </article>

        
           
            <div className="flex items-center py-8">
                <a href="#" className="h-10 w-10 bg-blue-800 hover:bg-blue-600 font-semibold text-white text-sm flex items-center justify-center">1</a>
                <a href="#" className="h-10 w-10 font-semibold text-gray-800 hover:bg-blue-600 hover:text-white text-sm flex items-center justify-center">2</a>
                <a href="#" className="h-10 w-10 font-semibold text-gray-800 hover:text-gray-900 text-sm flex items-center justify-center ml-3">التالي <i className="fas fa-arrow-right ml-2"></i></a>
            </div>

        </section>
   

            </Loyout>
        </>
    )
}
