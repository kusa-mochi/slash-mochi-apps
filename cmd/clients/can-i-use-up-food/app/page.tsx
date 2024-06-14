import TitleBar from "@/components/TitleBar/TitleBar";
import MainPane from "@/components/MainPane/MainPane";
import Image from "next/image";

export default function Home() {
  return (
    <main className="flex flex-col items-center justify-start h-screen">
      <TitleBar>Project Name</TitleBar>
      <MainPane></MainPane>
    </main>
    </ProjectContext.Provider>
  );
}
