"use client"

import TitleBar from "@/components/TitleBar/TitleBar";
import MainPane from "@/components/MainPane/MainPane";
import Image from "next/image";
import { ProjectContext, ProjectContextType } from "@/context/ProjectContext";
import { useContext, useState } from "react";

export default function Home() {
  const ctx = useContext(ProjectContext)
  return (
    <ProjectContext.Provider value={ctx}>
      <main className="flex flex-col items-center justify-start h-screen">
        <TitleBar>Project Name</TitleBar>
        <MainPane></MainPane>
      </main>
    </ProjectContext.Provider>
  );
}
