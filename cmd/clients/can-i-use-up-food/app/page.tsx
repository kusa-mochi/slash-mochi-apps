"use client"

import TitleBar from "@/components/TitleBar/TitleBar";
import MainPane from "@/components/MainPane/MainPane";
import Image from "next/image";
import { ProjectSettingProvider } from "@/context/ProjectSettingReducer";

export default function Home() {
  return (
    <ProjectSettingProvider>
      <main className="flex flex-col items-center justify-start h-screen">
        <TitleBar></TitleBar>
        <MainPane></MainPane>
      </main>
    </ProjectSettingProvider>
  );
}
