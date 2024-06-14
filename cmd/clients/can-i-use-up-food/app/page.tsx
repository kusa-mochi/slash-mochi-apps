"use client"

import TitleBar from "@/components/TitleBar/TitleBar";
import MainPane from "@/components/MainPane/MainPane";
import Image from "next/image";
import { ProjectContext } from "@/context/ProjectContext";
import { useState } from "react";

export default function Home() {
  const [test, setTest] = useState("ほげほげ")
  return (
    <ProjectContext.Provider value={test}>
      <main className="flex flex-col items-center justify-start h-screen">
        <TitleBar>Project Name</TitleBar>
        <MainPane></MainPane>
      </main>
    </ProjectContext.Provider>
  );
}
