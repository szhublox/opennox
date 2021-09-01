#ifndef NOX_PORT_GAME2
#define NOX_PORT_GAME2

#include "defs.h"


CHAR* nox_xxx_spriteDefByAlphabetAdd_44CD10(CHAR* a1);
int nox_xxx_keyFirstLetterNumberCli_44CD30(CHAR* a1);
char* nox_get_thing_name(int i);
nox_thing* nox_get_thing(int i);
wchar_t* nox_get_thing_pretty_name(int i);
wchar_t* nox_get_thing_desc(int i);
int nox_get_thing_pretty_image(int i);
int nox_xxx_getTTByNameSpriteMB_44CFC0(CHAR* a1);
int nox_xxx_getvalByName2Imp_44CFD0(CHAR* a1);
int sub_44D020(const void* a1, const void* a2);
int sub_44D040(int i);
BOOL sub_44D060(int a1);
int sub_44D090(int a1);
int nox_drawable_link_thing(nox_drawable* a1, int i);
void sub_44D2C0();
nox_thing* nox_get_things_head();
int sub_44D320(int a1);
int sub_44D330(CHAR* a1);
int sub_44D340(CHAR* a1);
void sub_44D3A0();
void sub_44D5C0(int a1, int a2);
int sub_44D640();
BOOL sub_44D660(const char* a1);
int sub_44D7E0(int a1);
void sub_44D8C0();
int sub_44D8F0();
int nox_xxx_playDialogFile_44D900(int a1, int a2);
int sub_44D930();
void sub_44D960();
int sub_44D970();
int sub_44D990();
int sub_44D9A0();
void sub_44D9D0();
int sub_44D9F0(int a1);
int sub_44DA60(int a1);
int sub_44DC40(int a1, int a2);
int sub_44DCA0(int a1, int a2);
BOOL sub_44DD00();
int nox_xxx_screenFadeEffect_44DD20(int a1);
int nox_client_drawFadingScreen_44DD70(int a1);
void nox_xxx_cliClearScreen_44DDC0();
int4* sub_44DDF0(int a1);
int4* sub_44DE30(int a1);
void sub_44DE80(_DWORD* a1);
void sub_44DEE0(_DWORD* a1);
struc_36* sub_44DF50();
void sub_44DF70();
int nox_client_setScreenFade_44DF90(int a1);
int sub_44DFD0(int a1);
void sub_44E000();
void sub_44E020();
void sub_44E040();
int sub_44E070();
int sub_44E090();
int nox_xxx_cliPlayMapIntro_44E0B0(int a1);
BOOL sub_44E0D0();
_DWORD* sub_44E110();
void sub_44E320();
int nox_client_getIntroScreenDuration_44E3B0();
void sub_44E3C0();
int nox_xxx_playGMCAPsmth_44E3E0();
_DWORD* sub_44E560();
int nox_client_wndQuestBriefProc_44E630(int a1, int a2, int a3, int a4);
BOOL nox_xxx_wndProc_44E6E0(int a1, int a2, int a3, int a4);
int sub_44E6F0(_DWORD* a1, int xLeft);
double sub_44E8B0();
int sub_44E8D0();
int nox_client_lockScreenBriefing_450160(int a1, int a2, char a3);
int sub_450560();
int sub_450570();
int sub_450580();
int sub_4505B0();
int sub_4505E0();
unsigned __int8 sub_450750();
char sub_450760(char a1);
int sub_450960(const void* a1, const void* a2);
char* sub_450AD0(char* a1);
int sub_450AF0(int a1);
void nox_gui_setQuestStage_450B00(int a1);
int nox_gui_getQuestStage_450B10();
void nox_xxx_unused_451430();
int nox_xxx_unused_4514F0();
int sub_451850(int a2, int a3);
int sub_451920(_DWORD* a2);
void sub_451970();
int sub_4519C0();
int sub_451BE0(int a1);
int sub_451CA0(_DWORD* a1);
int sub_451F30(int a1, int a2);
int sub_451F90(int a1);
int sub_451FE0(int a1);
int sub_452010();
void sub_452050(_DWORD* a1);
int* sub_452120(int a1);
_DWORD** sub_452190(int a1);
int* sub_4521A0(int a1);
int sub_4521F0();
int***** sub_452230();
char* nox_xxx_draw_452270(int a1);
int sub_4522A0(int a1);
int sub_4522C0(int a1);
char* sub_4522E0(int a1);
_DWORD* nox_xxx_draw_452300(_DWORD* a1);
int sub_4523D0(_DWORD* a1);
int sub_452410(int a1);
int sub_452490(_DWORD* a1);
void sub_452510(int a3);
__int64 sub_452690(int a3, __int64 a4, int a5);
int sub_4526D0(int a1);
int sub_4526F0(int a1);
int* sub_452810(int a1, char a2);
BOOL nox_thing_read_AVNT_452890(int a1, void* a2);
BOOL nox_thing_skip_AVNT_452B00(nox_memfile* f);
BOOL nox_thing_skip_AVNT_inner_452B30(nox_memfile* f);
int sub_452BD0(int a1, char* a2);
void nox_xxx_clientPlaySoundSpecial_452D80(int a1, int a2);
void sub_452DC0(int a1, int a2, int a3);
void sub_452E10(int a1, int a2, int a3);
_DWORD* sub_452E60(_DWORD* a1);
_DWORD* sub_452E80(_DWORD* a1);
int sub_452E90(_DWORD* a1, int a2);
int sub_452EB0(int* a1);
int sub_452EE0(int a1, int a2);
unsigned int sub_452F10(int a1, int a2);
int sub_452F50(int a1, int a2);
_DWORD* sub_452F80(int a1, int a2);
int sub_452FA0(int a1);
int sub_452FE0(int a1, int a2);
int sub_453000(_DWORD* a1, unsigned int a2, int a3);
void sub_453050();
void nox_xxx____setargv_9_453060();
int sub_453070();
BOOL sub_453080(char a1);
_DWORD* sub_4532E0();
int sub_453350(int a1, int a2);
int sub_4533D0(int a1, int a2, int a3, int a4);
int* sub_4535E0(int* a1);
int sub_4535F0(int a1);
char* sub_453600();
int sub_453610();
_BYTE* sub_453620(_BYTE* a1, char a2, int a3);
_DWORD* sub_453640(_DWORD* a1, int a2, int a3);
BOOL sub_453660(int a1);
BOOL sub_453690(int a1);
int sub_4536B0(_DWORD* a1);
int sub_453710();
char sub_453750();
int sub_4537F0();
_DWORD* sub_453B00();
int sub_453B80(int a1, int a2);
void sub_453F70(const void* a1);
char* sub_453F90();
int sub_453FA0(int a1, int a2, int a3);
BOOL sub_454000(int a1, int a2);
int sub_454040(_DWORD* a1);
int sub_4540E0(int a1);
char sub_454120();
int nox_xxx_guiServerAccessLoad_4541D0(int a1);
int sub_454A90(int a1, int a2);
int nox_xxx_windowAccessProc_454BA0(int a1, int a2, int* a3, int a4);
int sub_455770();
int sub_4557D0(int a1);
int* sub_455800();
int sub_455920(int a1);
void sub_455950(wchar_t* a1);
int sub_4559B0(wchar_t* a1);
int sub_455A00(int a1);
char sub_455A50(char a1);
int sub_455C10();
int sub_455CD0(_BYTE* a1, _DWORD* a2);
int sub_455E70(unsigned __int8 a1);
int sub_455EE0();
int sub_455F10(int a1);
int sub_455F60();
int sub_456050();
int sub_456070();
int sub_4560D0(int a1, int a2);
int sub_456240();
int sub_456500();
int sub_456640(int a1, int a2);
char sub_456BB0(int a1);
wchar_t* sub_456D00(int a1, wchar_t* a2);
int* sub_456D60(int a1);
int sub_456DF0(int a1);
int sub_456E40(int a1, int a2);
int sub_456EA0(wchar_t* a1);
int sub_456F10(wchar_t* a1, int a2);
int sub_456FA0();
unsigned __int8 sub_457120(int a1);
int sub_457140(int a1, wchar_t* a2);
int sub_4571A0(int a1, int a2);
int* sub_457350(unsigned __int8 a1, unsigned __int8 a2);
int sub_4573A0();
void sub_4573B0();
int sub_457460(int a1);
int nox_xxx_guiServerOptionsTryHide_4574D0();
int sub_457B60(int a1);
int sub_457F30(int a1);
int nox_xxx_windowServerOptionsDrawProc_458500(_DWORD* a1, int a2);
int nox_xxx_guiServerOptionsProc_458590(int a1, int a2, int a3, int a4);
char sub_459150();
int sub_4593B0(int a1);
int sub_459560(int a1);
int sub_459700();
int* nox_xxx_guiServerOptionsHide_4597E0(int a1);
char* sub_459870();
int sub_459A40(char* a1);
char* sub_459AA0(int a1);
int sub_459C30();
int sub_459D50(int a1);
int sub_459D60();
int sub_459D70();
int sub_459D80(int a1);
BOOL sub_459DA0();
BOOL sub_459DB0(nox_drawable* dr);
void sub_459DD0(nox_drawable* dr, char a2);
int nox_xxx_cliRemoveHealthbar_459E30(nox_drawable* dr, char a2);
int nox_xxx_cliFirstMinimapObj_459EB0();
int nox_xxx_cliNextMinimapObj_459EC0(int a1);
int sub_459F00(nox_drawable* dr);
int sub_459F40_drawable(int a1);
_DWORD* sub_459F70(nox_drawable* dr);
nox_drawable* nox_xxx_cliGetSpritePlayer_45A000();
nox_drawable* sub_45A010(nox_drawable* dr);
int nox_get_drawable_count();
int nox_xxx_sprite_45A030();
int sub_45A040(int a1);
int sub_45A060();
int sub_45A070(int a1);
int sub_45A090();
int sub_45A0A0(int a1);
int sub_45A0C0();
int sub_45A0E0(int a1);
_DWORD* nox_xxx_sprite_45A110_drawable(nox_drawable* a1);
void sub_45A160_drawable(nox_drawable* a1);
BOOL nox_alloc_drawable_init(int cnt);
void nox_drawable_free();
nox_drawable* nox_xxx_spriteLoadAdd_45A360_drawable(int thingInd, int a2, int a3);
void nox_xxx_sprite_45A480_drawable(int a1);
int nox_xxx_spriteDelete_45A4B0(nox_drawable* dr);
void nox_xxx_spriteDeleteStatic_45A4E0_drawable(nox_drawable* dr);
void nox_xxx_spriteDeleteAll_45A5E0(int a1);
int nox_xxx_spriteIsPlayerSprite_45A630(int a1);
void sub_45A670(unsigned int a1);
_DWORD* nox_xxx_netSpriteByCodeDynamic_45A6F0(int a1);
_DWORD* nox_xxx_netSpriteByCodeStatic_45A720(int a1);
int nox_xxx_unused_45A750();
int sub_45A840(_DWORD* a1);
int nox_xxx_spriteSetActiveMB_45A990_drawable(int a1);
int nox_xxx_cliDestroyObj_45A9A0(int a1);
int* sub_45A9B0(int a1, int a2);
int sub_45AB40();
int nox_xxx_spriteSetFrameMB_45AB80(int a1, int a2);
int nox_xxx_guiSpellSortFn_45ABC0(const void* a1, const void* a2);
int nox_xxx_bookSetColor_45AC40();
void nox_client_toggleSpellbook_45AC70();
int nox_xxx_bookHideMB_45ACA0(int a1);
int nox_xxx_guiSpellSortList_45ADF0(int a1);
void nox_xxx_book_45B010(int a1);
int nox_xxx_bookWndProc_45B070(int a1, int a2);
int nox_xxx_bookClickSpell_45B1F0();
int nox_xxx_bookClickCreature_45B200();
int nox_xxx_book_45B210(int a1, int a2);
int nox_xxx_bookChildWndProcMB_45B360(_DWORD* a1, unsigned int a2);
int nox_xxx_bookListWndProc_45B5F0(int a1, unsigned int a2, unsigned int a3);
int nox_xxx_bookMoveToPage_45B930(int a1);
int nox_xxx_bookInit_45B9D0();
int nox_xxx_bookDrawIconFn_45CB30(_DWORD* a1);
int nox_xxx_bookWndFn_45CC10(_DWORD* a1, int a2, unsigned int a3);
int sub_45CFA0();
int sub_45CFC0();
void nox_xxx_netSpellRewardCli_45CFE0(int a1, int a2, int a3, int a4);
void nox_xxx_netGuideRewardCli_45D140(int a1, int a2);
int* nox_xxx_bookSetForward_45D200(int* a1, int a2, int2* a3);
void nox_xxx_abilityReward_45D290(int a1, char* a2, int a3);
int sub_45D320(int a1);
int sub_45D400(int a1);
char* nox_xxx_clientQuestDisableAbility_45D4A0(int a1);
int sub_45D500(int a1);
_DWORD* sub_45D550(_DWORD* a1);
void nox_xxx_bookFillAll_45D570(int a1, int a2);
int sub_45D7D0(int* a1, int* a2);
void sub_45D810();
int sub_45D9B0();
BOOL nox_xxx_guiSpellTest_45D9C0();
int nox_xxx_guiSpellTargetClickSet_45D9D0(int a1);
int nox_xxx_guiSpell_45DA10(int a1);
void nox_client_invokeSpellSlot_45DA50(int a1);
void nox_xxx_clientStoreLastButton_45DAD0(int a1);
int nox_xxx_clientSendAbil_45DAF0(int a1);
int nox_xxx_clientSendSpell_45DB20(char* a1, int a2, char a3);
int sub_45DB90();
void nox_xxx_guiSpellTargetClickCheckSend_45DBB0();
void* nox_xxx_book_45DBE0(void* a1, int a2, int a3);
void nox_xxx_spellKeyPackSetSpell_45DC40(int a1, int a2, int a3);
int nox_xxx_bookSpellDrop_45DCA0(int a1, char a2, int a3, int a4, int* a5);
int nox_xxx_updateSpellIcons_45DDF0(int a1);
int nox_xxx_updateSpellIconDir_45DE10(int a1, int a2);
int nox_xxx_spellBoxPointToWnd_45DE60(int a1, int a2, int a3);
int nox_xxx_guiSpellSetCursor_45DF60(int a1, char a2);
int sub_45DFC0(int a1);
int nox_xxx_clientUpdateButtonRow_45E110(int a1);
int nox_xxx_buttonsGetSelectedRow_45E180();
int nox_xxx_quickbarButtonBookDraw_45EF30();
int sub_45EF40();
int nox_xxx_quickBarWnd_45EF50(int a1, int a2, unsigned int a3);
void nox_xxx_clientSwapQuickbarKeys_45F300(int a1, int a2, int a3);
void sub_45F390(int a1, int a2, int a3, int a4);
int nox_xxx_quickbarButtonBookWnd_45F450(int a1, unsigned int a2);
int sub_45F500(int a1, int a2);
int sub_45F520(int a1, unsigned int a2);
int sub_45F5D0(_DWORD* a1);
int nox_xxx_quickbarTrapUpDownProc_45F630(int a1, unsigned int a2);
int nox_xxx_quickbarTrapUpDownDraw_45F6F0(_DWORD* a1);
int nox_xxx_quickbarTrapButtonProc_45F7A0(_DWORD* a1, unsigned int a2, unsigned int a3);
int sub_45F890();
BOOL nox_xxx_quickbar_45F8D0(int a1, int a2, int a3, int a4);
int sub_45F8F0();
void sub_45F900();
int nox_xxx_quickbarTrapProc_45FB90(int a1, unsigned int a2, int a3, int a4);
int nox_xxx_quickbarDrawFn_460000();
int nox_xxx_quickBarInitWindow_4601F0(int a1, int a2, int a3, int a4, int a5, int a6, int a7);
void* sub_4602F0();
void* sub_460380();
void nox_client_trapSetNext_4603A0();
void nox_client_trapSetPrev_4603F0();
void sub_460440();
int nox_client_trapSetSelect_4604B0(int a1);
int sub_4604E0();
int nox_client_spellSetNext_4604F0();
int nox_client_spellSetPrev_460540();
void nox_client_spellSetSelect_460590();
void sub_460630();
int nox_xxx_guiSpell_460650();
int sub_460660();
int nox_xxx_quickBarClose_4606B0();
void sub_460920();
int __thiscall sub_460940(void* this);
int sub_460A10(int a1, int a2, int a3, char a4);
int sub_460B90(int a1);

int nox_client_screenFadeTimeout_44DAB0(int a1, int a2, void (*a3)(void));
int nox_client_screenFadeXxx_44DB30(int a1, int a2, void (*a3)(void));
int sub_44DBA0(int a1, int a2, void (*a3)(void));
int sub_44DBF0(int a1, int a2, void (*a3)(void));

#endif // NOX_PORT_GAME2