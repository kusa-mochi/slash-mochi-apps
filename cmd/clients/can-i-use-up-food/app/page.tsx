"use client"

import TitleBar from "@/components/TitleBar/TitleBar";
import MainPane from "@/components/MainPane/MainPane";
import Image from "next/image";
import { ProjectSettingProvider } from "@/context/ProjectSettingReducer";
import { FoodsProvider } from "@/context/FoodsReducer";
import { CalendarItemsProvider } from "@/context/CalendarItemsReducer";
import { CalendarItemGroupsProvider } from "@/context/CalendarItemGroupsReducer";

export default function Home() {
  return (
    <ProjectSettingProvider>
      <FoodsProvider>
        <CalendarItemsProvider>
          <CalendarItemGroupsProvider>
            <main className="flex flex-col items-center justify-start h-screen text-gray-800">
              <TitleBar></TitleBar>
              <MainPane></MainPane>
            </main>
          </CalendarItemGroupsProvider>
        </CalendarItemsProvider>
      </FoodsProvider>
    </ProjectSettingProvider>
  );
}
